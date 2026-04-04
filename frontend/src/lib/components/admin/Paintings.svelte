<div
    onpointermove={(e) => {
        if (isDragging) {
            e.preventDefault()
            cursorY = e.clientY
            for (let i = 0; i < paintingElements.length; i++) {
                // don't forget that positions are 1 indexed!!!
                if (i == sourcePosition - 1) {
                    // don't do this on painting you're dragging
                    continue
                }
                const element = paintingElements[i].getBoundingClientRect()
                if (element.top < -100) {
                    // skip paintings that are outside the view
                    continue
                }
                if (i == 0 && element.top > cursorY) {
                    // if cursor is higher than paintings, put target on top
                    targetPosition = 1
                    break
                }
                if (element.top < cursorY && element.bottom > cursorY) {
                    const middle = element.top + (element.bottom - element.top) / 2
                    if (cursorY < middle) {
                        if (targetPosition !== i) {
                            targetPosition = i + 1
                        }
                    } else {
                        targetPosition = i + 2
                    }
                }
            }
            if (
                cursorY >
                paintingElements[paintingElements.length - 1].getBoundingClientRect().bottom
            ) {
                targetPosition = paintingElements.length + 1
            }
        }
    }}
    onpointerup={() => {
        if (isDragging) {
            movePainting()
            isDragging = false
            sourcePosition = 0
            targetPosition = 0
            shadowHeightOffset = 0
            shadowHeight = 0
        }
    }}
    class="flex min-h-screen flex-col items-center"
>
    <div class="flex w-225 flex-col">
        {#each data.paintings as painting, i}
            <div
                hidden={painting.position !== targetPosition}
                class="my-1.5 w-full items-center justify-center bg-gray-300 p-0.5"
                style="height: {shadowHeight}px;"
            ></div>
            <!-- if hovering over this painting, move it up. So position must == target -->
            <div
                bind:this={paintingElements[i]}
                class="my-1.5 flex w-full items-center gap-4 border bg-white p-0.5"
                style={sourcePosition == painting.position
                    ? "transform: translateY(" +
                      (cursorY - shadowHeightOffset) +
                      "px); position: fixed; top: 0; margin: 0; opacity: 0.85; width: 900px;"
                    : ""}
            >
                <svg
                    class="h-40 w-20 text-gray-600 hover:text-gray-800"
                    onpointerdown={(e) => {
                        if (e.button !== 0) return
                        e.preventDefault()
                        isDragging = true
                        sourcePosition = painting.position
                        targetPosition = painting.position
                        shadowHeight = paintingElements[i].offsetHeight
                        shadowHeightOffset = shadowHeight / 2
                        cursorY = e.clientY
                    }}
                    viewBox="0 0 16 16"
                >
                    <circle cx="6" cy="2" r="1.8" fill="currentColor" />
                    <circle cx="14" cy="2" r="1.8" fill="currentColor" />
                    <circle cx="6" cy="8" r="1.8" fill="currentColor" />
                    <circle cx="14" cy="8" r="1.8" fill="currentColor" />
                    <circle cx="6" cy="14" r="1.8" fill="currentColor" />
                    <circle cx="14" cy="14" r="1.8" fill="currentColor" />
                </svg>
                <img
                    src={"/images/" + painting.img_url}
                    alt={painting.name}
                    class="h-48 w-48 object-cover object-center"
                />
                <div class="flex flex-col justify-center gap-1 pt-1">
                    <p>Name: {painting.name}</p>
                    <p>Author: {painting.author}</p>
                    <p>Size: {painting.size}</p>
                    <p>Price: {painting.sold ? "n/a" : painting.price}</p>
                    <p>Technique: {painting.technique}</p>
                    <div class="flex gap-1">
                        <p>Description:</p>
                        <p>{painting.description}</p>
                    </div>
                    <div class="flex flex-col">
                        <p>This Painting:</p>
                        {#if painting.sold}
                            <p style="color: green;">was sold!</p>
                        {/if}
                        {#if painting.copiable}
                            <p style="color: green;">can be copied!</p>
                        {/if}
                        {#if painting.printable}
                            <p style="color: green;">can be printed!</p>
                        {/if}
                        {#if painting.favorite}
                            <p style="color: green;">is in your favorites</p>
                        {:else}
                            <p style="color: red;">is not in your favorites</p>
                        {/if}
                    </div>
                </div>
                <div class="grow"></div>
                <div class="mr-5 flex items-center gap-3">
                    <button
                        onclick={() => showEditModal(painting)}
                        class="min-w-20 cursor-pointer rounded-xl bg-blue-400 p-3">Edit</button
                    >
                    <button
                        onclick={() => deletePainting(painting.uuid)}
                        class="min-w-20 cursor-pointer rounded-xl bg-red-400 p-3">Delete</button
                    >
                </div>
            </div>
        {/each}
        <div
            hidden={paintingElements.length + 1 !== targetPosition}
            class="my-1.5 w-full items-center justify-center bg-amber-500 p-0.5"
            style="height: {shadowHeight}px;"
        ></div>
        <button
            onclick={() => (showModal = true)}
            class="mt-6 mb-10 w-40 cursor-pointer self-center rounded-xl bg-green-300 p-3"
            >Upload Painting</button
        >
    </div>
</div>
