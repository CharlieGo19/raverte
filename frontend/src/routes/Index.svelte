<script>
    import {Router, Route} from "svelte-routing";
    //import Home from "./Home.svelte.old"; // chart-only
    import Login from "./Login.svelte";

    export let url = "";
    let chartOnlyMode = false;

</script>

<Router {url}>
  
        <!--{#if chartOnlyMode }
        <Route path="/" component={Home} /> This would be login? Can we do conditional here 
        via svelte to check if chartOnlyMode then set route to either or login? 
        {:else}
        <Route path="/" component={Login} />
        {/if}-->
        <Route path="/" component={ Login } />
</Router>

