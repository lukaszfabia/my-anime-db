import { VoiceActor, Character } from "@/types/models";

export function createName<T extends VoiceActor | Character>(t: T | null): string {
    return t ? `${t.lastname} ${t.name}` : "";
}
