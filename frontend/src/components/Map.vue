<template>
  <div class="map-container">
    <l-map :options="{zoomControl: false, attributionControl: false}" ref="map" v-model:zoom="zoom" :use-global-leaflet="false" :center="center" @ready="onMapReady">
      <l-tile-layer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        layer-type="base"
        name="OpenStreetMap"
      ></l-tile-layer>
      <l-polyline v-if="routePolyline" :lat-lngs="routePolyline" />
    </l-map>
  </div>
</template>

<script>
import "leaflet/dist/leaflet.css";
import AutoComplete from 'primevue/autocomplete';
import Button from 'primevue/button';
import { LMap, LTileLayer, LPolyline } from "@vue-leaflet/vue-leaflet"; // Import LPolyline for displaying the route

export default {
  components: {
    LMap,
    LTileLayer,
    LPolyline, // Add the polyline component
    AutoComplete,
    Button,
  },
  data() {
    return {
      zoom: 12, // Adjust the zoom level
      center: [64, 17], // Center of the map
      routePolyline: null, // To store route coordinates as polyline
    };
  },
  methods: {
  onMapReady(map) {
      console.log('Map is ready:', map);
      // You can now use the map instance for further operations
    },
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

