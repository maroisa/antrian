import { For } from "solid-js";
import { loketData } from "../utils/socket";

export default function LoketMini() {
    return (
        <div
            class="grid grid-cols-2 md:min-w-md font-bold text-2xl md:text-4xl h-1/4 md:h-full"
            id="listLoket"
        >
            <For each={Object.entries(loketData())}>
                {(item) => (
                    <div
                        class={`flex justify-center items-center border-2 border-b-0 md:border-l-0 bg-${item[0]}`}
                    >
                        {item[0]}
                        {item[1]}
                    </div>
                )}
            </For>
        </div>
    );
}
