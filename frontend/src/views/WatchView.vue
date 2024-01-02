<script setup lang="ts">
import { useRoute } from "vue-router";
import toWebVTT from "srt-webvtt";
import { onMounted, onUnmounted } from "vue";
import { GetTorrentInfo, StartDownload, StopDownload } from "../../wailsjs/go/main/App";
import BackButton from "@/components/BackButton.vue";

const route = useRoute();
const infoHash = typeof route.params.infoHash == "string" ? route.params.infoHash : route.params.infoHash[0];
const fileName = typeof route.params.fileName == "string" ? route.params.fileName : route.params.fileName[0];

const videoSrc = encodeURI(`/stream/${infoHash}/file/${fileName}`);

onMounted(async () => {
  console.log(`Starting download infoHash: ${infoHash}, fileName: ${fileName}`);
  await StartDownload(infoHash, fileName);

  console.log("Starting download subtitles");
  const info = await GetTorrentInfo(infoHash);
  const subtitleFiles = info.files.filter((file) => file.displayPath.endsWith(".srt"));
  let i = 0;
  for (const file of subtitleFiles) {
    console.log(`Downloading subtitle ${file.displayPath}`);
    const textTrackUrl = await normalizeSubtitleURL(file.displayPath);
    appendSubtitle(file.displayPath, textTrackUrl, i === 0);
    i++;
  }
});

onUnmounted(async () => {
  console.log(`Stopping download infoHash: ${infoHash}, fileName: ${fileName}`);
  await StopDownload(infoHash, fileName);
});

const normalizeSubtitleURL = async (fileName: string): Promise<string> => {
  const resp = await fetch(encodeURI(`/stream/${infoHash}/file/${fileName}`));
  const data = await resp.blob();
  return await toWebVTT(data);
};

const appendSubtitle = (label: string, textTrackUrl: string, isDefault: boolean) => {
  const textTrack = document.createElement("track");
  textTrack.kind = "subtitles";
  textTrack.label = label;
  textTrack.srclang = "en";
  textTrack.src = textTrackUrl;
  textTrack.default = isDefault;
  const video = document.querySelector("video");
  video?.appendChild(textTrack);
};


</script>

<template>
  <main class="w-full">
    <div class="flex justify-between m-4">
      <h1 class="text-xl ">{{ fileName }}</h1>
      <BackButton />
    </div>
    <video controls class="w-full">
      <source
        :src="videoSrc"
        type="video/mp4"
      />
    </video>
  </main>
</template>