import { A } from "@solidjs/router";
import { For } from "solid-js";
import { ChevronLeft } from "../components/Svg";
import { loketData } from "../utils/socket";

export default function Loket() {
    return (
        <main class="flex flex-col justify-center items-center p-10 ">
            <div class="w-full">
                <A href=".." class="btn w-fit">
                    <ChevronLeft className="size-4" />
                    Kembali
                </A>
            </div>
            <div
                class="h-full w-full grid grid-cols-2 max-w-6xl gap-4 bg-white p-8 rounded-lg shadow-lg"
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
