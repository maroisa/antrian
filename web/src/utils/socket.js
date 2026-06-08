import { createSignal } from "solid-js";

const wsProtocols = window.location.protocol == "http:" ? "ws:" : "wss:";
const ws = new WebSocket(
    wsProtocols + "//" + import.meta.env.VITE_API_HOST + "/loket/ws",
);
const [loketData, setLoketData] = createSignal({
    A: 0,
    B: 0,
    C: 0,
    D: 0,
    E: 0,
    F: 0,
});
const [newLoket, setNewLoket] = createSignal(null);

ws.onopen = () => console.log("ws connected");
ws.onclose = () => console.log("ws disconnected");

ws.onmessage = (event) => {
    const changedKey = Object.keys(event.data).find(
        (key) => loketData()[key] !== event.data[key],
    );
    const result = changedKey ? `${changedKey}${event.data[changedKey]}` : "";
    setNewLoket(result);

    setLoketData(JSON.parse(event.data));
};

export { ws as socket, loketData, newLoket };
