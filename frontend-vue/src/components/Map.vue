<template>
  <div class="map-container">
    <l-map ref="map" v-model:zoom="zoom" :use-global-leaflet="false" :center="[64, 17]">
      <l-tile-layer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        layer-type="base"
        name="OpenStreetMap"
      ></l-tile-layer>
    </l-map>
    <Button label="Submit" @click="fetchRoute" />
  </div>
</template>

<script>
import "leaflet/dist/leaflet.css";
import AutoComplete from 'primevue/autocomplete';
import Button from 'primevue/button';
import { LMap, LTileLayer } from "@vue-leaflet/vue-leaflet";
import { getRoute } from '@/services/api.js';

export default {
  components: {
    LMap,
    LTileLayer,
    AutoComplete,
    Button,
  },
  methods: {
    async fetchRoute() {
      try {
        const response = await getRoute(1);
        console.log('Route data:', response.data);
      } catch (error) {
        console.error('Error fetching route:', error);
      }
    },
  },
  data() {
    return {
      zoom: 5,
    };
  },
};
</script>

<style scoped>
.map-container {
  width: 100%; /* Ensure full width */
  height: 80vh; /* Height set to 80% of the viewport height */
  padding: 20px; /* Padding around the map */
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1); /* Subtle shadow */
  border-radius: 10px; /* Rounded corners */
  margin: 0 auto; /* Center horizontally */
  background-color: #fff; /* Background color */
}

l-map {
  width: 100%; /* Set width of the map */
  height: 100%; /* Set height of the map */
}
</style>

