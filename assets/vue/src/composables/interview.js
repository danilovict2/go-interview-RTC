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
                isLoading.value = false;
            })
            .catch((err) => toast.error(err));
    };

    return { getInterviews };
}
