<script setup lang="ts">
import { ref } from "vue";
import toWebVTT from "srt-webvtt";

const props = defineProps<{
  title: string;
  src: string;
  subtitles?: { label: string; url: string; }[];
}>();

const videoRef = ref<HTMLVideoElement | null>(null);
const subtitles = ref<{ label: string; url: string; }[]>(props.subtitles ?? []);


function removeAddSubtitles() {
  for (const track of videoRef.value?.querySelectorAll("track") ?? []) videoRef.value?.removeChild(track);
}

async function changeSubtitle(label: string, url: string, srclang?: string) {
  console.log(`Changing subtitle to ${label} at ${url}`);
  removeAddSubtitles();

  const resp = await fetch(url);
  const textTrackUrl = await toWebVTT(await resp.blob());

  const textTrack = document.createElement("track");
  textTrack.kind = "subtitles";
  textTrack.label = label;
  textTrack.srclang = srclang ?? "vn";
  textTrack.src = textTrackUrl;
  textTrack.default = true;
  videoRef.value?.appendChild(textTrack);
}

</script>

<template>
  <section>
    <video
      ref="videoRef"
      controls autoplay disablepictureinpicture preload crossorigin="anonymous"
      class="w-full my-2"
      :title="$props.title">
      <source :src="$props.src" />
    </video>
    <div class="subtitle-controller mt-4">
      <h3 class="text-lg my-4"><span class="border-b-2 border-red-700 pb-1">Subtitles</span></h3>
      <div class="-mx-1">
        <template v-for="sub in subtitles" :key="sub.url">
          <button
            class="bg-stone-900 hover:bg-stone-800 text-slate-100 px-4 py-2 m-1 rounded"
            @click="changeSubtitle(sub.label, sub.url)"
          >
            {{ sub.label }}
          </button>
        </template>
      </div>
    </div>
  </section>
</template>