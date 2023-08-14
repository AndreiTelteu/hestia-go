import { onMount } from "solid-js";

export default function GuestLogin(props) {
    
    onMount(() => {
        props?.windowUpdateProps({
            title: 'Login',
        });
    });
    
    return <div>
        login form
    </div>
}
