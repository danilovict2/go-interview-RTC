<template>
    <Card
        class="group relative overflow-hidden hover:border-primary/50 transition-all duration-300 hover:shadow-lg cursor-pointer"
    >
        <div
            :class="`absolute inset-0 bg-gradient-to-br ${action.gradient} opacity-100 group-hover:opacity-50
            transition-opacity`"
        />

        <div class="relative p-6 size-full">
            <div class="space-y-3">
                <div
                    :class="`w-12 h-12 rounded-full flex items-center justify-center bg-${action.color}/10
                    group-hover:scale-110 transition-transform`"
                >
                    <component :is="action.icon" :class="`size-6 text-${action.color}`"></component>
                </div>

                <div class="space-y-1">
                    <h3 class="font-semibold text-xl group-hover:text-primary transition-colors">
                        {{ action.title }}
                    </h3>
                    <p class="text-sm text-muted-foreground">{{ action.description }}</p>
                </div>
            </div>
        </div>
    </Card>
</template>

<script setup>
import Card from './ui/card/Card.vue';

defineProps({
    action: Object,
});
</script>
