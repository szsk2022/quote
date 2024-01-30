# HTML页面使用SZSK一言API

#### 在任意位置添加script代码
```html
 <script>        document.addEventListener("DOMContentLoaded", function () {
            getRandomQuote();

            function getRandomQuote() {
                fetch('http://localhost:8080/?lang=cn') //这里改为你搭建的API网址，如果要请求英语，则改为en
                    .then(response => response.json())
                    .then(data => {
                        if (data.error) {
                            document.getElementById('quote').innerText = '加载失败，请稍候重试';
                        } else {
                            document.getElementById('quote').innerText = data.quote;
                        }
                    })
                    .catch(error => console.error('Error:', error));
            }
        });
    </script>
```
### 在合适位置调用API

```html
<div id="quote">正在加载...</div>
```

### 使用官网版API  
>每隔5s自动刷新

```html
  <script src="https://www.sunzishaokao.com/cdn/yy.js"></script>
  <link rel="stylesheet" href="https://www.sunzishaokao.com/cdn/yy.css"> // 按需添加
  <div id="quote">正在加载...</div>
'''
