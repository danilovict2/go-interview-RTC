<template>
    <DropdownMenu>
        <DropdownMenuTrigger as-child>
            <Button variant="ghost" size="sm">
                <img class="size-8 rounded-full"
                    src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                    alt="" v-if="authStore.authUser">
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
    Cookies.remove('jwt')
    authStore.$reset();

    router.push({ name: 'login' })
}
</script>