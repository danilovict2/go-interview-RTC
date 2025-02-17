<template>
    <AppLoading v-if="isLoading"></AppLoading>
    <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8" v-else>
        <div class="sm:mx-auto sm:w-full sm:max-w-sm">
            <h2 class="mt-10 text-center text-2xl/9 font-bold tracking-tight">
                Create a new account
            </h2>
        </div>
        <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            <form class="space-y-6" @submit="onSubmit">
                <FormField name="first_name" v-slot="{ componentField }">
                    <FormItem v-auto-animate>
                        <FormLabel class="block text-sm/6 font-medium">First Name</FormLabel>
                        <FormControl class="mt-2">
                            <Input
                                type="text"
                                v-bind="componentField"
                                placeholder="First name..."
                                class="block w-full rounded-md px-3 py-1.5 text-base outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                            />
                        </FormControl>
                        <FormMessage />
                    </FormItem>
                </FormField>
                <FormField name="last_name" v-slot="{ componentField }">
                    <FormItem v-auto-animate>
                        <FormLabel class="block text-sm/6 font-medium">Last Name</FormLabel>
                        <FormControl class="mt-2">
                            <Input
                                type="text"
                                v-bind="componentField"
                                placeholder="Last name..."
                                class="block w-full rounded-md px-3 py-1.5 text-base outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                            />
                        </FormControl>
                        <FormMessage />
                    </FormItem>
                </FormField>
                <FormField name="email" v-slot="{ componentField }">
                    <FormItem v-auto-animate>
                        <FormLabel class="block text-sm/6 font-medium">Email</FormLabel>
                        <FormControl class="mt-2">
                            <Input
                                type="email"
                                v-bind="componentField"
                                placeholder="Enter your email..."
                                class="block w-full rounded-md px-3 py-1.5 text-base outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                            />
                        </FormControl>
                        <FormMessage />
                    </FormItem>
                </FormField>
                <FormField name="password" v-slot="{ componentField }">
                    <FormItem v-auto-animate>
                        <FormLabel class="block text-sm/6 font-medium">Password</FormLabel>
                        <FormControl class="mt-2">
                            <Input
                                type="password"
                                v-bind="componentField"
                                placeholder="Create a password..."
                                class="block w-full rounded-md px-3 py-1.5 text-base outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                            />
                        </FormControl>
                        <FormMessage />
                    </FormItem>
                </FormField>

                <Button
                    type="submit"
                    class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                    Register
                </Button>
            </form>

            <p class="mt-10 text-center text-sm/6 text-gray-500">
                Already have an account?
                <RouterLink to="/login" class="font-semibold text-indigo-600 hover:text-indigo-500">
                    Sign In!
                </RouterLink>
            </p>
        </div>
    </div>
</template>

<script setup>
import { useForm } from 'vee-validate';
import { toTypedSchema } from '@vee-validate/zod';
import * as z from 'zod';

import { Button } from '@/components/ui/button';
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import axios from 'axios';
import { toast } from 'vue3-toastify';
import router from '@/router';
import AppLoading from '@/components/AppLoading.vue';
import { ref } from 'vue';

const formSchema = toTypedSchema(
    z.object({
        first_name: z.string().min(1).max(50),
        last_name: z.string().min(1).max(50),
        email: z.string().min(2).max(50).email(),
        password: z.string().min(8).max(50),
    }),
);

const { handleSubmit } = useForm({
    validationSchema: formSchema,
});

const isLoading = ref(false);

const onSubmit = handleSubmit((values) => {
    isLoading.value = true;
    axios
        .post('/users/store', values, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        })
        .then((resp) => {
            document.cookie = `jwt=${resp.data.token};expires=${resp.data.expires};path=/;secure;`;
            isLoading.value = false;
            router.push({ name: 'Home' });
        })
        .catch((err) => toast.error(err.response.data));
});
</script>
