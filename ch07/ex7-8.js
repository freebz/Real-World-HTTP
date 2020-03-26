// 예제 7-8 데이터 채널 초기화

var connection = new RTCPeerConnection();
var dataChannel =
  connection.createDataChannel("data channel", {
    ordered: false,       // 순서 보증?
    maxRetransmitTime: 0, // 계속 재전송하는 기간
    maxRetransmits: 0	  // 재전송 횟수
  });
