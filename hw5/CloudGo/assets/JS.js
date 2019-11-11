/*
    Author: 颜屹豪
    Group: Night-Nine
    Extra:设置了多个难度，稍微美化了一下界面，图片是使用多年的壁纸(2333, 希望你能喜欢。
*/

/*常量区*/
let plen = 200;                     //方块的长宽
let num = 25;                       //方块的数量
let sqnum = Math.sqrt(num);        //数量开方
let staset = [0,1,2,3,4,5,6,7,8];   //初始位置
let isstart = false;               //游戏是否开始
let pset;                           //当前状态
let place = num-1;                 //空方块定位
let steps;                         //步数
let bgimg = 'url("img/background.jpg")';

/*预加载*/
$(function () {
    $("#button").click(restart);
    $("#display").val('请选择要挑战的拼图数，再点击开始游戏');
    $(".circle").click(clcrad);
});

/*控制器*/
function game() {
    if(isstart===true) {
        let ts = $(".pict").index(this);
        if(ts>=0&&ts!==place) {
            let sub = pset[ts] - pset[place];
            if (sub <= sqnum&&sub>=-sqnum) {
                if (sub % sqnum === 0) {
                    movpic(ts);
                }
                else if (pset[place] % sqnum === 0) {
                    if (sub===1)
                        movpic(ts);
                }
                else if (pset[place] % sqnum === sqnum-1) {
                    if (sub===-1)
                        movpic(ts);
                }
                else if (sub === 1 || sub === -1) {
                    movpic(ts);
                }
            }
            if (pset.toString() === staset.toString())
                gameOver();//游戏结束
        }
    }
}
function float() {
    $(this).animate({
        'opacity': '0.8',
        'transform': 'scale(1.2,1.2)',
        'z-index': '20'
    },50);
}

function down() {
    $(this).animate({
        'opacity': '1',
        'transform': 'scale(0.8,0.8)',
        'z-index': '0'
    },50);
}

/*移动方块*/
function movpic(ts) {
    /*取对象*/
    let st = $(".pict").eq(ts);
    let ed = $(".pict").eq(place);
    /*取位置（没有直接使用对应参数是为了解决因为多次点击导致位置失真的情况）*/
    let stlestr = (plen*(pset[ts]%sqnum)+'px');
    let sttopstr = (plen*(Math.floor(pset[ts]/sqnum))+'px');
    let edlestr = (plen*(pset[place]%sqnum)+'px');
    let edtopstr = (plen*(Math.floor(pset[place]/sqnum))+'px');
    /*移动*/
    st.animate({
        "left":edlestr,
        "top":edtopstr
    },500);
    ed.css({
        "left":stlestr,
        "top":sttopstr
    });
    let temp = pset[ts];
    pset[ts] = pset[place];
    pset[place] = temp;
    steps++;
    $('#display').val('Your steps: '+steps);
}

/*游戏结束*/
function gameOver() {
    isstart = false;
    alert('Congratulations!');
    $("#display").val('恭喜你还原了拼图, 你共走了'+steps+'步');
}
/*重设画面*/
function reset() {
    isstart = false;
    num = parseInt($('input[name="cset"]:checked').val());
    sqnum = Math.sqrt(num);
    place = num-1;
    plen = 600/sqnum;
    let gamer = $("#game");
    gamer.css({
        'background-image': 'unset'
    });
    gamer.empty();
    for(let temp=0;temp<num;temp++) {
        gamer.append("<div></div>");
        gamer.children("div:last-child").addClass("pict");
        let lestr = (plen*(temp%sqnum)+'px');
        let topstr = (plen*(Math.floor(temp/sqnum))+'px');
        let balestr = (-plen*(temp%sqnum)+'px');
        let batopstr = (-plen*(Math.floor(temp/sqnum))+'px');
        if (temp!==num-1) {
            $(".pict").eq(temp).css({
                'border-color': 'skyblue',
                'border-style': 'solid',
                'border-width': '1px',
                'position': 'absolute',
                'width': plen+'px',
                'height': plen+'px',
                "left": lestr,
                "top": topstr,
                'background-image': bgimg,
                "background-position-x": balestr,
                "background-position-y": batopstr
            });
        }
        else {
            $(".pict").eq(temp).css({
                "left": lestr,
                "top": topstr,
                "opacity": '0'
            });
        }
    }
    $(".pict").click(game);
    $(".pict").mouseenter(float);
    $(".pict").mouseleave(down);
}
/*点击选项*/
function clcrad() {
    $('#reset').html('开始游戏');
    reset();
    switch (num) {
        case 9:
            $('#display').val('拿这个练练手吧，请点击开始游戏');
            break;
        case 16:
            $('#display').val('我觉得这个有点难，不过开始游戏吧');
            break;
        case 25:
            $('#display').val('看起来你有点膨胀，开始游戏吧!');
            break;
        case 100:
            $('#display').val('说真的，我怕你承受不来！');
    }
}

/*重新开始*/
function restart() {
    if($('input[name="cset"]:checked').length!==0) {
        $('#reset').html('重新开始');
        reset();
        let rd;
        do {
            pset = [];
            for (let temp = 0; temp < num; temp++) {
                do {
                    rd = Math.floor(Math.random() * num);
                } while (pset.indexOf(rd) !== -1);
                pset.push(rd);
            }
        } while (!isvalid(pset));
        for (let temp = 0; temp < num; temp++) {
            let lestr = (plen * (pset[temp] % sqnum) + 'px');
            let topstr = (plen * (Math.floor(pset[temp] / sqnum)) + 'px');
            $(".pict").eq(temp).animate({
                "left": lestr,
                "top": topstr
            });
        }
        steps = 0;
        $('#display').val('Your steps: '+steps);
        isstart = true;
    }
}

/*拼图是否可还原*/
function isvalid(arr) {
    let judge = (sqnum+sqnum)%2;
    let count = 0;
    for(let temp1=0;temp1<num-1;temp1++) {
        for(let temp2=temp1+1;temp2<num;temp2++) {
            if(arr[temp1]>arr[temp2])
                count++;
        }
    }
    count+=arr[place]%sqnum+Math.floor(arr[place]/sqnum);
    count%=2;
    if (count===judge)
        return true;
    return false;
}