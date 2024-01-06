import "./style.css";

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
import { NotifyAlert } from "../wailsjs/go/main/App";

const app = createApp(App);

app.config.errorHandler = async (err, vm, info) => {
  console.log("Error: ", err);
  console.log("Error Info: ", info);

  await NotifyAlert("Error", JSON.parse(JSON.stringify(err)));
}

app.use(router);

library.add(faPlay, faPause, faVolumeUp, faVolumeMute, faExpand, faCompress, faSpinner);

app.component("font-awesome-icon", FontAwesomeIcon);

app.mount("#app");
