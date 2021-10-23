import 'tailwindcss/tailwind.css'
import type { AppProps } from 'next/app'
import Register from '.'

export default function MyApp({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />
}