<template>
    <DropdownMenu>
        <DropdownMenuTrigger as-child>
            <Button variant="ghost" size="sm">
                <img
                    v-if="authStore.authUser"
                    :src="`https://api.dicebear.com/9.x/initials/svg/seed=${authStore.authUser.first_name}-${authStore.authUser.last_name}`"
                    alt="avatar"
                    class="size-8 rounded-full"
                />
                <UserIcon v-else />
            </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent class="w-56">
            <DropdownMenuItem @click="logout" v-if="authStore.authUser" class="cursor-pointer">
                <LogOut class="mr-2 h-4 w-4" />
                <span>Log out</span>
            </DropdownMenuItem>
            <div v-else>
                <RouterLink to="/login">
                    <DropdownMenuItem class="cursor-pointer">
                        <LogIn class="mr-2 h-4 w-4" />
                        <span>Sign In</span>
                    </DropdownMenuItem>
                </RouterLink>

                <RouterLink to="/register">
                    <DropdownMenuItem class="cursor-pointer">
                        <UserPlus class="mr-2 h-4 w-4" />
                        <span>Create an account</span>
                    </DropdownMenuItem>
                </RouterLink>
            </div>
        </DropdownMenuContent>
    </DropdownMenu>
</template>

<script setup>
import { Button } from '@/components/ui/button';

import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import router from '@/router';
import { useAuthStore } from '@/stores/auth';
import Cookies from 'js-cookie';

import { LogIn, LogOut, UserIcon, UserPlus } from 'lucide-vue-next';

const authStore = useAuthStore();

const logout = () => {
    Cookies.remove('jwt');
    authStore.$reset();

    router.push({ name: 'login' });
};
</script>
