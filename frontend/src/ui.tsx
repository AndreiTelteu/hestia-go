export {css, createGlobalStyles} from 'solid-styled-components'
import {css, type StylesArg} from 'solid-styled-components'

export type Stylesheet<T> = {
    [key in keyof T]: StylesArg
}
export type StylesheetReturn<T> = {
    [key in keyof T]: string
}

export type StylesheetFn<T> = (stylesheet: Stylesheet<T>) => StylesheetReturn<T>

export function stylesheet<T>(stylesheet: Stylesheet<T>): StylesheetReturn<T> {
    Object.keys(stylesheet).forEach(key => stylesheet[key] = css(stylesheet[key]))
    return stylesheet as StylesheetReturn<T>
}
