// 예제 10-2 CORS가 이면에서 이루어지는 통신 예

fetch('https://api.external.com', {
      method: 'PATCH',
      mode: 'cors',
}).then(function (response) {
      // 요청이 지날 때 호출되는 콜백
});
