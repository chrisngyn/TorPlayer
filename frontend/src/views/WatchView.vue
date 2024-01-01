<script setup lang="ts">
import {
  CaptionControl,
  Captions,
  ClickToPlay,
  ControlGroup,
  Controls,
  ControlSpacer,
  DefaultSettings,
  FullscreenControl,
  PlaybackControl,
  Player,
  ScrubberControl,
  SettingsControl,
  Spinner,
  TimeProgress,
  Ui,
  Video,
  VolumeControl,
} from "@vime/vue-next";
import { useRoute } from "vue-router";
import { onMounted, onUnmounted } from "vue";
import { StartDownload, StopDownload } from "../../wailsjs/go/main/App";

const route = useRoute();
const infoHash = typeof route.params.infoHash == "string" ? route.params.infoHash : route.params.infoHash[0];
const fileName = typeof route.params.fileName == "string" ? route.params.fileName : route.params.fileName[0];

const videoSrc = encodeURI(`/stream/${infoHash}/file/${fileName}`);

onMounted(async () => {
  console.log(`Starting download infoHash: ${infoHash}, fileName: ${fileName}`);
  await StartDownload(infoHash, fileName);
});

onUnmounted(async () => {
  console.log(`Stopping download infoHash: ${infoHash}, fileName: ${fileName}`);
  await StopDownload(infoHash, fileName);
});

</script>

<template>
  <main class="w-full">
    <h1 class="m-4 text-xl ">{{fileName}}</h1>
    <Player style="--vm-player-theme: #dc2626;">
      <Video>
        <source
          :src="videoSrc"
          type="video/mp4"
        />
      </Video>
      <Ui>
        <ClickToPlay />
        <Captions />
        <Spinner />
        <DefaultSettings />
        <Controls hide-on-mouse-leave full-width :active-duration="2000" pin="bottomLeft">
          <ControlGroup>
            <ScrubberControl />
          </ControlGroup>
          <ControlGroup space="top">
            <PlaybackControl tooltip-direction="right" />
            <VolumeControl />
            <TimeProgress />
            <ControlSpacer />
            <CaptionControl />
            <SettingsControl />
            <FullscreenControl keys="f" tooltip-direction="left" />
          </ControlGroup>
        </Controls>
      </Ui>
    </Player>
  </main>
</template>