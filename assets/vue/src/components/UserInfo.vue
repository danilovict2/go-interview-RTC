<template>
    <div class="flex items-center gap-2">
        <Avatar class="h-6 w-6">
            <AvatarImage
                :src="`https://api.dicebear.com/9.x/initials/svg/seed=${user.first_name}-${user.last_name}`"
            />
            <AvatarFallback>
                <UserCircle class="h-4 w-4" />
            </AvatarFallback>
        </Avatar>
        <span>{{ user.first_name + ' ' + user.last_name }}</span>
    </div>
</template>

<script setup>
import { UserCircle } from 'lucide-vue-next';
import Avatar from './ui/avatar/Avatar.vue';
import AvatarFallback from './ui/avatar/AvatarFallback.vue';
import AvatarImage from './ui/avatar/AvatarImage.vue';

const { user } = defineProps({
    user: Object,
});
</script>
