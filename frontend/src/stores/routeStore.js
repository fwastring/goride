import { defineStore } from 'pinia';
import { ref } from 'vue';
import { deleteRoute, getRoutes, createRoute } from '@/services/api'; // Import your API functions

export const useRouteStore = defineStore('route', {
  state: () => ({
    routes: [], // Stores the list of routes
    rider_start_address: "", // Rider's start address
    rider_end_address: "", // Rider's end address
  }),
  
  actions: {
    // Adds a new route to the routes array
    async addRoute(route) {
		const response = await createRoute(route)	
		await this.fetchRoutes()	
    },

    // Sets the rider's start and end addresses
    setRiderAddresses(start, end) {
      this.rider_start_address = start;
      this.rider_end_address = end;
    },

    // Fetches all routes from the API
    async fetchRoutes() {
      try {
        const response = await getRoutes();
        this.routes = response.data;
      } catch (error) {
        console.error('Error fetching routes:', error);
      }
    },

    // Sets routes directly, useful when updating routes after an operation
    setRoutes(routes) {
      this.routes = routes;
    },

    // Filters routes by provided IDs, leaving only those that match the route IDs
    async filterRoutes(routeIDs) {
		if (routeIDs == null) {
			this.routes = {}
		} else {
			const response = await getRoutes();
			this.routes = response.data.filter(route => routeIDs.includes(route.id));
		}
    },

    // Removes a route by its ID from both the API and the state
    async removeRoute(routeId) {
      try {
        await deleteRoute({ id: routeId.toString() });
        this.routes = this.routes.filter(route => route.id !== routeId);
      } catch (error) {
        console.error('Error deleting route:', error);
      }
    },
  }
});

