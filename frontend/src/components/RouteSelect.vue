<template>
  <div class="route-container">
    <div v-for="route in routes" :key="route.ID" class="route-card">
      <Card style="width: 25rem; overflow: hidden">
        <template #title>{{ route.StartAddress }} - {{ route.EndAddress }}</template>
        <template #content>
          <p class="m-0">
            {{ route.details }}
          </p>
        </template>
        <template #footer>
          <div class="flex gap-4 mt-1">
            <Button label="Delete" severity="secondary" outlined class="w-full" @click="handleCancel(route.ID)" />
            <Button label="Display" class="w-full" @click="displayRoute(route)" />
          </div>
        </template>
      </Card>
    </div>
  </div>
</template>
<script>
import { ref, onMounted } from 'vue';
import Button from 'primevue/button';
import Card from 'primevue/card';
import { getRoutes} from '@/services/api.js';

export default {
  components: {
    Button,
    Card,
  },
  setup(props, { emit }) {
    const routes = ref([]);

    const fetchRoute = async () => {
      try {
        const response = await getRoutes();
        console.log(response.data);
        routes.value = response.data; // Update reactive `routes` array
      } catch (error) {
        console.error('Error fetching route:', error);
      }
    };

    const handleCancel = async (routeId) => {
      try {
        // await removeRoute(routeId); // Call API to remove route
        // routes.value = routes.value.filter((route) => route.id !== routeId); // Remove route from the list
      } catch (error) {
        console.error('Error removing route:', error);
      }
    };

    const displayRoute = (route) => {
      emit('display-route', route); // Emit event to parent to handle route display
    };

    onMounted(() => {
      fetchRoute(); // Fetch routes when the component is mounted
    });

    return {
      routes,
      fetchRoute,
      handleCancel,
      displayRoute,
    };
  },
};
</script>

<style scoped>
/* Add any styles if needed */
</style>

