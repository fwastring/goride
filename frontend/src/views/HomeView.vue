<script>
import RouteSelect from '../components/RouteSelect.vue';
import RouteCreator from '../components/RouteCreator.vue';
import RouteSearcher from '../components/RouteSearcher.vue';
import Map from '../components/Map.vue';

export default {
  components: {
    RouteSelect,
    RouteCreator,
    RouteSearcher,
    Map,
  },
  data() {
    return {
      selectedRoute: null, // Store the route to be displayed on the map
    };
  },
  methods: {
    handleDisplayRoute(route) {
      this.selectedRoute = route; // Set the selected route to be passed to the map
    },
  },
};
</script>

<template>
  <main class="layout-container">
	<div class="side-container">
	  <RouteCreator />
	  <RouteSearcher />
      <RouteSelect @display-route="handleDisplayRoute" />
    </div>
    <Map class="map-container" :route="selectedRoute" />
  </main>
</template>

<style>
.main-container {
  display: flex;
  justify-content: space-between; /* Ensure space between the containers */
  padding: 1rem;
  gap: 1rem;
}

.side-container {
  flex: 1; /* Take up less space */
  max-width: 300px; /* Limit the width of side containers */
}

.map-container {
  flex: 5; /* Map takes up more space */
  height: 80vh; /* Keep height as per your design */
}
</style>
