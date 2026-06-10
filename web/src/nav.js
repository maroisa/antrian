import { SpeakerWaveIcon } from "@heroicons/vue/24/outline";
import { ComputerDesktopIcon } from "@heroicons/vue/24/outline";
import { PlusCircleIcon } from "@heroicons/vue/24/outline";

const NAV = [
    {
        to: "/panggil",
        name: "Panggil",
        color: "primary",
        icon: SpeakerWaveIcon,
    },
    {
        to: "/daftar",
        name: "Daftar",
        color: "neutral",
        icon: PlusCircleIcon,
    },
    {
        to: "/view",
        name: "Tampilan",
        color: "secondary",
        icon: ComputerDesktopIcon,
    },
];
export default NAV;
