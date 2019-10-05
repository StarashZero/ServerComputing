package main

import (
	"fmt"
	"os"
	"github.com/spf13/pflag"
	"io"
	"bufio"
	"os/exec"
)

/*
	selpg_args:  参数结构体
*/
type selpg_args struct {
	start_page  int			//开始页
	end_page    int			//结束页
	in_filename string		//输入文件
	page_len    int			//页长度
	page_type   bool		//是否按页结束符计算(默认为按页长度计算)
	print_dest  string		//打印机地址
}

var progname string			//程序名

const INT_MAX = int(^uint(0) >> 1);

func usage() {
	fmt.Fprintf(os.Stderr, "\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n", progname)
}

//处理程序参数
func (sa *selpg_args)process_args() {
	//绑定各参数
	pflag.IntVarP(&sa.start_page, "start_page", "s", 0, "Start page")
	pflag.IntVarP(&sa.end_page, "end_page", "e", 0, "End page")
	pflag.BoolVarP(&sa.page_type, "page_type", "f", false, "Page type")
	pflag.IntVarP(&sa.page_len, "page_len", "l", 72, "Lines per page")
	pflag.StringVarP(&sa.print_dest, "dest", "d", "", "Destination")
	pflag.Usage = func() {
		usage()
		pflag.PrintDefaults()
	}
	pflag.Parse()
	sa.in_filename = ""
	if remain := pflag.Args(); len(remain) > 0 {
		sa.in_filename = remain[0]
	}

	//判断各参数是否合法
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		pflag.Usage()
		os.Exit(1)
	}
	if sa.start_page < 1 || sa.start_page > (INT_MAX-1) {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %d\n", progname, sa.start_page)
		pflag.Usage()
		os.Exit(2)
	}
	if sa.end_page < 1 || sa.end_page > (INT_MAX-1) || sa.end_page < sa.start_page {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %d\n", progname, sa.start_page)
		pflag.Usage()
		os.Exit(3)
	}
	if sa.page_len < 1 || sa.page_len > (INT_MAX-1) {
		fmt.Fprintf(os.Stderr, "%s: invalid page length %d\n", progname, sa.page_len)
		pflag.Usage()
		os.Exit(4)
	}
	if sa.in_filename != "" {
		if _, err := os.Stat(sa.in_filename); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n", progname, sa.in_filename)
			pflag.Usage()
			os.Exit(5)
		}
	}
}

//运行逻辑
func (sa selpg_args)process_input() {
	var reader *bufio.Reader		//输入读取
	var writer io.WriteCloser		//输出写入

	//获得reader
	if sa.in_filename == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		fin, err := os.Open(sa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n", progname, sa.in_filename)
			os.Exit(6)
		}
		reader = bufio.NewReader(fin)
		defer fin.Close()
	}

	//获得writer
	if sa.print_dest == "" {
		writer = os.Stdout
	} else {
		cmd := exec.Command("lp","-d"+ sa.print_dest)
		var err error
		if writer, err = cmd.StdinPipe(); err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open pipe to \"%s\"\n",
				progname, sa.print_dest)
			fmt.Println(err)
			os.Exit(7)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "%s: cmd start error\n",
				progname)
			fmt.Println(err)
			os.Exit(8)
		}
	}

	line_ctr, page_ctr, pLen := 1, 1, sa.page_len
	ptFlag := '\n'
	if sa.page_type {
		ptFlag = '\f'
		pLen = 1
	}

	//使用reader读取选中页的数据并写入writer
	for {
		line, crc := reader.ReadString(byte(ptFlag));
		if crc != nil && len(line) == 0 {
			break
		}
		if line_ctr > pLen {
			page_ctr++
			line_ctr = 1
		}
		if page_ctr >= sa.start_page && page_ctr <= sa.end_page {
			_, err := writer.Write([]byte(line))
			if err != nil {
				fmt.Println(err)
				os.Exit(9)
			}
		}
		line_ctr++
	}

	//判断读取是否成功或者是否完成
	if page_ctr < sa.start_page {
		fmt.Fprintf(os.Stderr,
			"\n%s: start_page (%d) greater than total pages (%d),"+
				" no output written\n", progname, sa.start_page, page_ctr)
	} else if page_ctr < sa.end_page {
		fmt.Fprintf(os.Stderr, "\n%s: end_page (%d) greater than total pages (%d),"+
			" less output than expected\n", progname, sa.end_page, page_ctr)
	}
}

func main() {
	sa := selpg_args{}
	progname = os.Args[0]
	sa.process_args()
	sa.process_input()
}
