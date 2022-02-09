package badge

import "strings"

var likeBadgeTemplate = strings.TrimSpace(`
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Width}}" height="{{.Height}}">
      <linearGradient id="smooth" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
        <stop offset="1" stop-opacity=".1"/>
      </linearGradient>
      <mask id="round">
        <rect width="{{.Width}}" height="{{.Height}}" rx="{{.XRadius}}" ry="{{.YRadius}}" fill="#fff"/>
      </mask>
      <g mask="url(#round)">
        <rect width="{{.React.Rect.Bound.Width}}" height="{{.React.Rect.Bound.Height}}" fill="{{.React.Rect.Color}}" fill-opacity="{{.Opacity}}"/>
        <rect x="{{.Count.Rect.Bound.X}}" width="{{.Count.Rect.Bound.Width}}" height="{{.Count.Rect.Bound.Height}}" fill="{{.Count.Rect.Color}}" fill-opacity="{{.Opacity}}"/>
        <rect x="{{.Share.Rect.Bound.X}}" width="{{.Share.Rect.Bound.Width}}" height="{{.Share.Rect.Bound.Height}}" fill="{{.Share.Rect.Color}}" fill-opacity="{{.Opacity}}"/>
        <rect width="{{.Width}}" height="{{.Height}}" fill="url(#smooth)" fill-opacity="{{.Opacity}}"/>
      </g>
      <g fill="$fff" text-anchor="middle" font-family="{{.FontFamily}}" font-size="{{.FontSize}}">
        <text x="{{.Count.Text.Bound.X}}" y="{{.Count.Text.Bound.Y}}" fill="{{.Count.Text.Color}}" fill-opacity=".3">{{.Count.Text.Msg | html}}</text>
        <text x="{{.Count.Text.Bound.X}}" y="{{.Count.Text.Bound.Y}}" fill="#000">{{.Count.Text.Msg | html}}</text>
      </g>
      <g class="react_icon" transform="translate({{.React.Icon.Bound.X}},{{.React.Icon.Bound.Y}})">
        <path transform="scale(0.035,0.035)" d="M 433.601 67.001 C 408.901 42.301 376.201 28.801 341.301 28.801 C 306.401 28.801 273.601 42.401 248.901 67.101 L 236.001 80.001 L 222.901 66.901 C 198.201 42.201 165.301 28.501 130.401 28.501 C 95.601 28.501 62.801 42.101 38.201 66.701 C 13.501 91.401 -0.099 124.201 0.001 159.101 C 0.001 194.001 13.701 226.701 38.401 251.401 L 226.201 439.201 C 228.801 441.801 232.301 443.201 235.701 443.201 C 239.101 443.201 242.601 441.901 245.201 439.301 L 433.401 251.801 C 458.101 227.101 471.701 194.301 471.701 159.401 C 471.801 124.501 458.301 91.701 433.601 67.001 Z"
        class="{{.ReactClassName}}"
        onclick="react()"
        fill-opacity="0" fill="{{.React.Icon.Color}}" stroke-width="20" stroke="{{.React.Icon.Color}}"/>
      </g>
      <g class="share_icon" transform="translate({{.Share.Icon.Bound.X}}, {{.Share.Icon.Bound.Y}})">
        <path transform="scale(0.57,0.57)" d="M 125.01 47.955 C 123.524 47.955 122.193 48.633 121.31 49.696 L 115.565 46.434 C 115.744 45.932 115.842 45.391 115.842 44.828 C 115.842 44.265 115.744 43.724 115.565 43.221 L 121.309 39.959 C 122.192 41.022 123.523 41.701 125.01 41.701 C 127.662 41.701 129.82 39.542 129.82 36.889 C 129.82 34.237 127.663 32.079 125.01 32.079 C 122.358 32.079 120.2 34.237 120.2 36.889 C 120.2 37.452 120.298 37.993 120.477 38.496 L 114.732 41.758 C 113.849 40.696 112.518 40.018 111.032 40.018 C 108.379 40.018 106.221 42.175 106.221 44.828 C 106.221 47.48 108.379 49.638 111.032 49.638 C 112.518 49.638 113.849 48.96 114.732 47.897 L 120.477 51.159 C 120.298 51.662 120.2 52.202 120.2 52.766 C 120.2 55.418 122.358 57.576 125.01 57.576 C 127.662 57.576 129.82 55.418 129.82 52.766 C 129.82 50.113 127.663 47.955 125.01 47.955 Z M 125.01 33.762 C 126.734 33.762 128.137 35.165 128.137 36.889 C 128.137 38.614 126.734 40.018 125.01 40.018 C 123.286 40.018 121.883 38.614 121.883 36.889 C 121.883 35.165 123.286 33.762 125.01 33.762 Z M 111.032 47.955 C 109.307 47.955 107.904 46.552 107.904 44.828 C 107.904 43.104 109.307 41.701 111.032 41.701 C 112.756 41.701 114.159 43.104 114.159 44.828 C 114.159 46.552 112.756 47.955 111.032 47.955 Z M 125.01 55.893 C 123.286 55.893 121.883 54.49 121.883 52.766 C 121.883 51.041 123.286 49.638 125.01 49.638 C 126.734 49.638 128.137 51.041 128.137 52.766 C 128.137 54.49 126.734 55.893 125.01 55.893 Z"
        fill="{{.Share.Icon.Color}}"/>
      </g>
      <style>
        .react_icon:hover{
          cursor:pointer;
        }
        .react_icon .like{
          fill-opacity:1;
        }
        .share_icon:hover{
          cursor:pointer;
        }
      </style>
    </svg>
`)
