// 예제 7-4 웹소켓의 접속과 데이터 전송

var socket = new WebSocket('ws://game.example.com:12010/updates');
socket.onopen = () => {
    setInterval(() => {
	if (socket.bufferedAmount === 0) {
	    socket.send(getUpdateData());
	}
    }, 50);
};
