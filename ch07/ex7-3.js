// 예제 7-3 EventSource 예제

const evtSource = new EventSource("ssedemo.php");

// 메시지의 이벤트 핸들러
evtSource.onmessage = (e) => {
    const newElement = document.createElement("li");
    newElement.innerHTML = "message: " + e.data;
    eventList.appendChild(newElement);
};

evtSource.addEventListener("ping", (e) => {
    const newElement = document.createElement("li");

    const obj = JSON.parse(e.data);
    newElement.innserHTML = "ping at " + obj.time;
    eventList.appendChild(newElement);
}, false);
