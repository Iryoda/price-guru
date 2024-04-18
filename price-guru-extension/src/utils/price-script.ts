export const handleSelectPrice = async () => {
    document.body.style.cursor = 'pointer'

    console.log('Test')

    let originalBg = ''

    const checkIfElementHasSearchableTag = (target: HTMLElement) => {
        if (target.id || target.className) return true

        return false
    }

    const makeBgPretty = (e: MouseEvent) => {
        const target = e.target as HTMLElement
        originalBg = target.style.background
        target.style.background = 'orange'
    }

    const makeBgOriginal = (e: MouseEvent) => {
        const target = e.target as HTMLElement
        target.style.background = originalBg
    }

    const onClick = async (e: MouseEvent) => {
        const target = e.target as HTMLElement
        target.style.background = originalBg

        const hasSearchableProps = checkIfElementHasSearchableTag(target)

        console.log(target, hasSearchableProps)

        document.removeEventListener('mouseover', makeBgPretty)
        document.removeEventListener('mouseout', makeBgOriginal)
        document.removeEventListener('click', onClick)

        chrome.runtime.sendMessage({
            message: 'price-clicked',
            target: target.outerHTML,
            location: window.location.href,
        })
    }

    document.addEventListener('mouseover', makeBgPretty)
    document.addEventListener('mouseout', makeBgOriginal)
    document.addEventListener('click', onClick)
}
