<template>
  <div class="map-container">
    <l-map ref="map" v-model:zoom="zoom" :use-global-leaflet="false" :center="center" @ready="onMapReady">
      <l-tile-layer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        layer-type="base"
        name="OpenStreetMap"
      ></l-tile-layer>
      <l-polyline v-if="routePolyline" :lat-lngs="routePolyline" />
    </l-map>
    <Button label="Fetch Route" @click="fetchRoute" />
  </div>
</template>

<script>
import "leaflet/dist/leaflet.css";
import AutoComplete from 'primevue/autocomplete';
import Button from 'primevue/button';
import { LMap, LTileLayer, LPolyline } from "@vue-leaflet/vue-leaflet"; // Import LPolyline for displaying the route
import { getRoute } from '@/services/api.js';

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
    async fetchRoute() {
      try {
        const response = await getRoute(8);
		  console.log(response.data.Geometry.coordinates)
        const routeData = response.data.Geometry.coordinates; // Assuming your API returns an array of coordinates
        // Example format for routeData: [{ lat: 64, lng: 17 }, { lat: 64.1, lng: 17.1 }, ...]
		  this.routePolyline = routeData.map(point => [point[1], point[0]])
        this.center = this.routePolyline[0]; // Optionally, center map on first point of route
        console.log('Route data:', this.routePolyline);
      } catch (error) {
        console.error('Error fetching route:', error);
      }
    },
    onMapReady() {
      console.log("Map is ready!");
    }
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

