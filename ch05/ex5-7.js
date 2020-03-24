// 예제 5-7 인터넷 익스플로러로 JSON을 획득

...
xhr.onload = function(e) {
    if (this.status == 200) {
	// IE11 대책
	var json = JSON.parse(xhr.responseText);
	/* IE 이외
        var json = xhr.response;
	*/
	console.log(josn)
    }
};
/* IE 이외
xhr.responseType = 'json';
*/
xhr.send();
