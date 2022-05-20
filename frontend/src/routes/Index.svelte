<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';

    const getChartOnlyMode = () => {
        window.go.raverte.Raverte.GetChartOnlyMode().then(
            (/** @type boolean */ res) => {
                return res
            }
        );
    }

    onMount(() => {
       if (getChartOnlyMode()) {
           goto('/mainstage');
       }else{
           goto('/login');
       }       
    });
</script>