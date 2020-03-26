// 예제 7-1 Fetch API 사용예

fetch("news.json", {
    method: 'GET',
    mode: 'cors',
    credentials: 'include',
    cache: 'default',
    headers: {
	'Content-Type': 'application/json'
    }
}).then((response) => {
    return response.json();
}).then((json) => {
    console.log(json);
});
