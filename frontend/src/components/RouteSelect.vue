<template>
  <div class="route-container">
    <Button label="Fetch Route" @click="fetchRoute" />
    <div v-for="route in routes" :key="route.ID" class="route-card">
      <Card style="width: 25rem; overflow: hidden">
        <template #title>{{ route.ID }}</template>
        <!-- <template #subtitle>{{ route.description }}</template> -->
        <!-- <template #content> -->
        <!--   <p class="m-0"> -->
        <!--     {{ route.details }} -->
        <!--   </p> -->
        <!-- </template> -->
        <template #footer>
          <div class="flex gap-4 mt-1">
            <Button label="Cancel" severity="secondary" outlined class="w-full" />
            <Button label="Save" class="w-full" />
          </div>
        </template>
      </Card>
    </div>
  </div>
</template>

<script>
import Button from 'primevue/button';
import Card from 'primevue/card';
import { getRoutes } from '@/services/api.js';

export default {
  components: {
    Button,
	Card,
  },
  data() {
    return {
	  routes: [],
    };
  },
  methods: {
    async fetchRoute() {
      try {
        const response = await getRoutes();
		  console.log(response.data)
		  this.routes = response.data
      } catch (error) {
        console.error('Error fetching route:', error);
      }
    },
  onMapReady(map) {
      console.log('Map is ready:');
      // You can now use the map instance for further operations
    },
	  mounted() {
    this.fetchRoutes(); // Fetch routes when the component is mounted
  },
	}
};
</script>

<style scoped>
</style>

