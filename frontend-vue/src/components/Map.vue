<template>
  <div style="height:800px; width:1000px">
    <l-map ref="map" v-model:zoom="zoom" :use-global-leaflet="false" :center="[64, 17]">
		<div class="input-container">
			<AutoComplete class="input-field" v-model="fromValue" :suggestions="items" @complete="search" />
			<AutoComplete class="input-field" v-model="tovalue" :suggestions="items" @complete="search" />
		</div>
      <l-tile-layer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        layer-type="base"
        name="OpenStreetMap"
      ></l-tile-layer>
    </l-map>
  </div>
</template>

<script>
import "leaflet/dist/leaflet.css";
import AutoComplete from 'primevue/autocomplete';
import { LMap, LTileLayer, LControl } from "@vue-leaflet/vue-leaflet";

export default {
  components: {
    LMap,
    LTileLayer,
	LControl,
	AutoComplete,
  },
  data() {
    return {
		zoom: 5,
			
    };
  },
};
</script>

<style>

.input-container {
	z-index: 400;
	border-radius: 25px;
	background: #73AD21;
	padding-left: 50px;
	padding-top: 20px;
	display: flex;
	flex-direction: column;
}
.input-field {
	z-index: 500;
	padding: 20px;
}
</style>
