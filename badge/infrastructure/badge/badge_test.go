package badge

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_Parsing_좋아요뱃지_성공(t *testing.T) {
	//given
	bg, border, icon, react, textColor, text, share, edge := "#fff", "black", "#333333", "#ff6767", "black", "1234", "#333333", "round"
	b := NewHeartBadge(bg, border, icon, react, false, textColor, text, share, edge)
	wr := HeartBadgeWriter
	expected := []byte(strings.TrimSpace(fmt.Sprintf(`
<html lang="en">
<head>
    <title>likeIt</title>
    <style>
        html{margin:0;}
        body{margin: 1px 0 0 2px;}
        .badge-container {display: block; background-color:%s;}
        .badge-container .badge-box {border: 2px solid %s;box-shadow: 0 0 1px 1px #bfbfbf; display: inline-block;padding: 5px 20px;border-radius: 25px; }
        .badge-container .badge-box .icon:hover {cursor:pointer;}
        .badge-container .badge-box .item-container {display: inline-block;}
        .badge-container .badge-box .hearts {padding-top: 6px;}
        .badge-container .badge-box .hearts .icon {transition-duration:0.3s;fill-opacity:0;fill:%s; stroke-width:35; stroke:%s;}
        .badge-container .badge-box .hearts .icon:hover{transition-duration:0.3s;stroke:%s;}
        .badge-container .badge-box .hearts .icon.react {transition-duration:0.3s; fill-opacity:1;stroke:%s;}
        .badge-container .badge-box .text .text-box{display: table;}
        .badge-container .badge-box .text .text-box span {display: table-cell;color:%s;vertical-align: middle; padding-bottom: 1px;}
        .badge-container .badge-box .share {margin-left: 20px;position: relative;}
        .tooltip-text {visibility: hidden;font-size: 0.8em;width: 50px;background-color: #333333;color: #fafafa;text-align: center;padding: 0 0 5px;border-radius: 6px;position: absolute;z-index: 1;top: 25px;}
    </style>
    <script src="https://cdn.jsdelivr.net/npm/clipboard@2/dist/clipboard.min.js"></script>
</head>
<body>
<div class="badge-container" >
    <div class="badge-box">
        <div class="hearts item-container" data-service="react" onclick="onHeartEvent(this)">
            <svg xmlns="http://www.w3.org/2000/svg" width="22px" height="21px">
                <g id="heart" class="icon " transform="translate(1,0)" >
                    <path transform="scale(0.04,0.04)"
                          d="M 433.601 67.001 C 408.901 42.301 376.201 28.801 341.301 28.801 C 306.401 28.801 273.601 42.401 248.901 67.101 L 236.001 80.001 L 222.901 66.901 C 198.201 42.201 165.301 28.501 130.401 28.501 C 95.601 28.501 62.801 42.101 38.201 66.701 C 13.501 91.401 -0.099 124.201 0.001 159.101 C 0.001 194.001 13.701 226.701 38.401 251.401 L 226.201 439.201 C 228.801 441.801 232.301 443.201 235.701 443.201 C 239.101 443.201 242.601 441.901 245.201 439.301 L 433.401 251.801 C 458.101 227.101 471.701 194.301 471.701 159.401 C 471.801 124.501 458.301 91.701 433.601 67.001 Z"
                    />
                </g>
            </svg>
            <div class="text item-container">
                <div class="text-box">
                    <span id="heart-text">%s</span>
                </div>
            </div>
        </div>
        <div id="share-url" class="share item-container" data-clipboard-text="url1">
            <div id="tooltip-text" class="tooltip-text">copy!</div>
            <svg class="icon" xmlns="http://www.w3.org/2000/svg" width="23" height="23" viewBox="0 0 24 24"
                 fill="none" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M16 4h2a2 2 0 012 2v4M8 4H6a2 2 0 00-2 2v14a2 2 0 002 2h12a2 2 0 002-2v-2"/>
                <rect x="8" y="2" width="8" height="4" rx="1" ry="1"/>
                <path d="M21 14H11"/>
                <path d="M15 10l-4 4 4 4"/>
            </svg>
        </div>
    </div>
</div>
<script>
    const clipboard = new ClipboardJS(document.getElementById('share-url'));
    clipboard.on('success', function (){
        document.getElementById('tooltip-text').style.visibility = 'visible';
        setTimeout(function (){
            document.getElementById('tooltip-text').style.visibility = 'hidden';
        },1000);
    })
    const onHeartEvent = function(e){
        let s = e.dataset.service;
        let cnt = document.getElementById('heart-text').innerText;

        if (s === 'react'){
            document.getElementById('heart').classList.add('react');
            e.dataset.service = '';
            document.getElementById('heart-text').innerText = String(Number(cnt) + 1);
        }else {
            document.getElementById("heart").classList.remove('react');
            e.dataset.service = 'react';
            document.getElementById('heart-text').innerText = String(Number(cnt) - 1);
        }
    }
</script>
</body>
</html>
`, bg, border, react, icon, react, react, textColor, text)))

	//when
	svg1, err := wr.ParseFile(*b)
	if err != nil {
		panic(err)
	}

	//then
	assert.Equal(t, svg1, expected)
}