<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="row items-center q-col-gutter-md">
              <div class="col-auto">
                <q-icon name="tournament" size="lg" class="text-primary" />
              </div>
              <div class="col">
                <div class="text-h4">Tournaments</div>
                <div class="text-subtitle1">Manage tournament brackets</div>
              </div>
              <div class="col-auto">
                <q-btn
                  color="primary"
                  label="Create Tournament"
                  icon="add"
                  @click="showCreateDialog = true"
                />
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="text-center q-pa-lg">
              <q-icon name="tournament" size="4rem" class="text-grey-5" />
              <div class="text-h6 text-grey-5 q-mt-md">Tournaments Coming Soon</div>
              <div class="text-body2 text-grey-5">Tournament management will be implemented in the next version</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>

    <q-dialog v-model="showCreateDialog">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">Create Tournament</div>
        </q-card-section>
        <q-card-section>
          <div class="text-body2">Tournament creation feature is coming soon!</div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn flat label="OK" color="primary" @click="showCreateDialog = false" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script setup>
import { ref } from 'vue'

const showCreateDialog = ref(false)
</script>