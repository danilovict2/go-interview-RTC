<template>
    <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
        <div class="sm:mx-auto sm:w-full sm:max-w-sm">
            <h2 class="mt-10 text-center text-2xl/9 font-bold tracking-tight">
                Sign in to your account
            </h2>
        </div>
        <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            <form class="space-y-6" @submit="onSubmit">
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
                                placeholder="Enter your password..."
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
                    Sign In
                </Button>
            </form>

            <p class="mt-10 text-center text-sm/6 text-gray-500">
                Don't have an account?
                <RouterLink
                    to="/register"
                    class="font-semibold text-indigo-600 hover:text-indigo-500"
                >
                    Create one for free!
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

const formSchema = toTypedSchema(
    z.object({
        email: z.string().min(2).max(50).email(),
        password: z.string().min(8).max(50),
    }),
);

const { handleSubmit } = useForm({
    validationSchema: formSchema,
});

const onSubmit = handleSubmit((values) => {
    axios
        .post('/login', values, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        })
        .then((resp) => {
            document.cookie = `jwt=${resp.data.token};expires=${resp.data.expires};path=/;secure;`;
            router.push({ name: 'home' });
        })
        .catch((err) => toast.error(err.response.data));
});
</script>
