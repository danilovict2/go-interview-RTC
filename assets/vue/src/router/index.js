import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import { useAuthStore } from '@/stores/auth';
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView,
            meta: {
                requiresAuth: true,
            },
        },
        {
            path: '/register',
            name: 'register',
            component: () => import('../views/RegisterView.vue'),
            meta: {
                requiresAuth: false,
            }
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('../views/LoginView.vue'),
            meta: {
                requiresAuth: false,
            }
        },
        {
            path: '/meetings/:id',
            name: 'meeting',
            component: () => import('../views/MeetingView.vue'),
            meta: {
                requiresAuth: true,
            },
        },
        {
            path: '/recordings',
            name: 'recordings',
            component: () => import('../views/RecordingsView.vue'),
            meta: {
                requiresAuth: true,
            },
        },
        {
            path: '/schedule',
            name: 'schedule',
            component: () => import('../views/ScheduleView.vue'),
            meta: {
                requiresAuth: true,
            },
        },
        {
            path: '/dashboard',
            name: 'dashboard',
            component: () => import('../views/DashboardView.vue'),
            meta: {
                requiresAuth: true,
            },
        },
        { 
            path: '/:pathMatch(.*)*', 
            name: 'NotFound', 
            component: () => import('../views/NotFoundView.vue') 
        },
    ],
});

router.beforeEach(async (to, from, next) => {
    const authStore = useAuthStore();
    await authStore.loadAuthUser();

    if (to.meta.requiresAuth && !authStore.authUser) {
        next('/login');
    } else if (to.meta.requiresAuth === false && authStore.authUser) {
        next('/');
    } else {
        next();
    }
});

export default router;
