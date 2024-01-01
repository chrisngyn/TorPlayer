import "./style.css";
import "@vime/core/themes/default.css";

import { createApp } from "vue";

import App from "./App.vue";
import router from "./router";
import { library } from "@fortawesome/fontawesome-svg-core";

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

import {
  faCompress,
  faExpand,
  faPause,
  faPlay,
  faSpinner,
  faVolumeMute,
  faVolumeUp,
} from "@fortawesome/free-solid-svg-icons";

const app = createApp(App);

app.use(router);

library.add(faPlay, faPause, faVolumeUp, faVolumeMute, faExpand, faCompress, faSpinner);

app.component("font-awesome-icon", FontAwesomeIcon);

app.mount("#app");
