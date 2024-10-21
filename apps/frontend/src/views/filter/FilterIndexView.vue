<template>
  <div>
    <h1>Filters</h1>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            Name
          </th>
          <th class="text-left">
            Description
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in filters"
          :key="item.name"
        >
          <td>{{ item.name }}</td>
          <td>{{ item.description }}</td>
        </tr>
      </tbody>
    </v-table>
  </div>
</template>

<script lang="ts" setup>
  import Filter from '../../lib/interfaces/filter.interface.ts';
  import { useFilterStore } from '../../store/filter.store.ts';

  useFilterStore().getFilters();

  const filters: Filter[] = useFilterStore().index;

</script>