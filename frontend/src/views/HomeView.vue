<script>
import RouteSelect from '../components/RouteSelect.vue';
import Map from '../components/Map.vue';

export default {
  components: {
    RouteSelect,
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
  <main>
    <Map :route="selectedRoute"/>
	<RouteSelect @display-route="handleDisplayRoute"/>
  </main>
</template>

