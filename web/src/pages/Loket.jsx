import { A } from "@solidjs/router";
import { createEffect, createSignal, For, on, onMount } from "solid-js";
import { ChevronLeft } from "../components/Svg";
import { loketData, newLoket } from "../utils/socket";

export default function Loket() {
    const [count, setCount] = createSignal(0);
    const [isSpeaking, setIsSpeaking] = createSignal(false);
    let synth = SpeechSynthesis;

    onMount(() => {
        if (typeof window !== "undefined" && "speechSynthesis" in window) {
            synth = window.speechSynthesis;
        }
    });

    createEffect(
        on(newLoket, (currentVal) => {
            if (count() < 2) {
                setCount(count() + 1);
                return;
            }
            speak(newLoket());
        }),
    );

    function speak(text) {
        if (!synth || !text) return;
        synth.cancel();
        const utterance = new SpeechSynthesisUtterance(text);
        const idVoice = synth
            .getVoices()
            .find((voice) => voice.lang.startsWith("id"));
        if (idVoice) {
            utterance.voice = idVoice;
        } else {
            utterance.lang = "id-ID";
        }

        utterance.onstart = () => setIsSpeaking(true);
        utterance.onend = () => setIsSpeaking(false);
        utterance.onerror = () => setIsSpeaking(false);

        console.log("Synth:", synth);
        console.log("Voice:", idVoice);
        console.log("Text:", text);
        synth.speak(utterance);
    }

    return (
        <main class="flex flex-col justify-center items-center p-10 ">
            <div class="w-full">
                <A href=".." class="btn w-fit">
                    <ChevronLeft className="size-4" />
                    Kembali
                </A>
            </div>
            <div
                class="h-full w-full grid grid-cols-2 max-w-6xl gap-4 bg-white p-8 rounded-lg shadow-lg border-4 "
                id="listLoket"
            >
                <For each={Object.entries(loketData())}>
                    {(item) => (
                        <div
                            class={`border-4 rounded-lg shadow-lg flex justify-center items-center font-bold text-6xl bg-${item[0]}`}
                        >
                            {item[0]}
                            {item[1]}
                        </div>
                    )}
                </For>
            </div>
        </main>
    );
}
