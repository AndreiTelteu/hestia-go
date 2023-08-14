import { For, lazy, JSX } from 'solid-js'
import { create100vh } from 'solidjs-div-100vh'
import { WindowManager } from 'solidjs-window-manager'
import { createGlobalStyles, css, stylesheet } from "./ui"


function App() {
  const fullHeight = create100vh();
  const GlobalStyle = createGlobalStyles`
    .window-manager-wrapper {
      height: ${() => fullHeight() ? fullHeight() + 'px' : '100vh'} !important;
      position: relative;
    }
  `
  
  let windowApi: { openWindow: (component: string, props?: any) => void };
  const pages: {
    page: string,
    label: string,
    icon: (props: JSX.ImgHTMLAttributes<{}>) => JSX.Element,
  }[] = [
    {
      page: 'Guest/Login',
      label: 'Login',
      icon: () => <p>icon</p>,
    }
  ]
  
  return (
    <div class={styles.root}>
      <GlobalStyle />
      <WindowManager
        loadWindow={(props) => lazy(() => import(/* @vite-ignore */ './pages/' + props.component))}
        onReady={(api) => (windowApi = api)}
        options={{
            persistent: true,
        }}
      >
        <div style={{ padding: '10px' }}>
          <div style={{ display: 'flex', 'flex-direction': 'column' }}>
            <For each={pages}>
              {(item) => (
                <div class={styles.desktopButtonContainer}>
                  <button
                      type="button"
                      class={styles.desktopButton}
                      onClick={() => windowApi?.openWindow?.(item.page)}
                  >
                    <item.icon class={styles.buttonIcon} />
                    <span class={styles.buttonText}>{item.label}</span>
                  </button>
                </div>
              )}
            </For>
          </div>
        </div>
      </WindowManager>
    </div>
  );
}

const styles = stylesheet({
  root: {
  },
  desktopButtonContainer: {
    marginBottom: '10px',
  },
  desktopButton: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    borderRadius: '5px',
    borderColor: 'transparent',
    backgroundColor: '#dae4eb71',
    width: '70px',
    transition: 'all 0.3s',
    cursor: 'pointer',
    '&:hover': {
      backgroundColor: '#dae4eb39',
    },
  },
  buttonIcon: {
    color: '#fff',
    height: '35px',
    width: '35px',
  },
  buttonText: {
    color: '#fff',
  },
})

export default App;
