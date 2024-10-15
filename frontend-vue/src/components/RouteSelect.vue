<template>
  <div class="route-container">
	  <Card style="width: 25rem; overflow: hidden">
		<template #title>Advanced Card</template>
		<template #subtitle>Card subtitle</template>
		<template #content>
			<p class="m-0">
				Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore sed consequuntur error repudiandae numquam deserunt quisquam repellat libero asperiores earum nam nobis, culpa ratione quam perferendis esse, cupiditate neque
				quas!
			</p>
		</template>
		<template #footer>
			<div class="flex gap-4 mt-1">
				<Button label="Cancel" severity="secondary" outlined class="w-full" />
				<Button label="Save" class="w-full" />
			</div>
		</template>
	</Card>
  </div>
</template>

<script>
import Button from 'primevue/button';
import Card from 'primevue/card';

export default {
  components: {
    Button,
	Card,
  },
  data() {
    return {
    };
  },
  methods: {
	}
};
</script>

<style scoped>
</style>

