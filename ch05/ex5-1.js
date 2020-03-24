// 예제 5-1 서버에서 이미지 파일을 가져와 img 태그로 표시

...
xhr.responseType = 'blob';
xhr.onload = function(e) {
    if (this.status == 200) {
        var blob = this.response;
	var img = document.createElement('img');
	img.onload = function(e) {
	    window.URL.revokeObjectURL(img.src);
	};
	img.src = window.URL.createObjectURL(blob);
	document.body.appendChild(img);
	...
    }
};
xhr.send();
