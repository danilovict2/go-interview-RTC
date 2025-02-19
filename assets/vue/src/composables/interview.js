import axios from 'axios';
import Cookies from 'js-cookie';
import { toast } from 'vue3-toastify';

export function useInterview() {
    const getInterviews = (dst, isLoading) => {
        isLoading.value = true;
        axios
            .get('/interviews', {
                headers: {
                    Authorization: `Bearer ${Cookies.get('jwt')}`,
                },
            })
            .then((resp) => {
                dst.value = resp.data.interviews;
            })
            .catch((err) => toast.error(err))
            .finally(() =>  isLoading.value = false);
    };

    return { getInterviews };
}
