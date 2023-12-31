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
  VolumeControl
} from '@vime/vue-next'
import {ref} from "vue";

const videoSrc = ref<string>(
  '/videos/video.mp4'
)
</script>

<template>
  <main class="w-full">
    <Player>
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