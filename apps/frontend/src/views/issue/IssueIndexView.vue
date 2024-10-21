<template>
  <div>
    <h1>Issues</h1>
    <v-table>
      <thead>
        <tr>
          <th class="text-left">
            Key
          </th>
          <th class="text-left">
            Summary
          </th>
          <th class="text-left">
            Status
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in issues"
          :key="item.key"
        >
          <td>{{ item.key }}</td>
          <td>{{ item.summary }}</td>
          <td>{{ item.status }}</td>
        </tr>
      </tbody>
    </v-table>
  </div>
</template>

<script lang="ts" setup>
  import Issue from '../../lib/interfaces/issue.interface.ts';
  import { useIssueStore } from '../../store/issue.store.ts';

  useIssueStore().getIssues();

  const issues: Issue[] = useIssueStore().index;

</script>
