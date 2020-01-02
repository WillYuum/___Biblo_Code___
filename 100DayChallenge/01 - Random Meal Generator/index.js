const MealInfo = []
const RandomMealApi = 'https://www.themealdb.com/api/json/v1/1/random.php'


const GetMealButton = document.getElementById("GetMealButton")
const Selectors = {
    header: document.querySelector(".MealInfo-header"),
    MealIngredient: document.querySelector(".MealIngredient"),
    vidContainer: document.querySelector(".vidContainer")
}




GetMealButton.addEventListener("click", async (event) => {
    event.preventDefault()
    let MealData = await getMealInfo()
    await CreateMealCard(MealData)
})


function CreateMealCard(meal) {

    const ingredients = [];
	// Get all ingredients from the object. Up to 20
	for(let i=1; i<=20; i++) {
		if(meal[`strIngredient${i}`]) {
			ingredients.push(`${meal[`strIngredient${i}`]} - ${meal[`strMeasure${i}`]}`)
		} else {
			// Stop if no more ingredients
			break;
		}
	}

    const HeaderMealSection = `
       <div class="MealImage">
           <img src=${meal.strMealThumb}
            alt="Displaying ${meal.strMeal}">
        </div>
    <div class="Header-right">
        <h2>${meal.strMeal}</h2>
    
        <p>${meal.strInstructions}</p>
    </div>`

    const MealIngredient = `
    ${meal.strTags ? `<p><strong>Tags:</strong> ${meal.strTags.split(',').join(', ')}</p>` : ''}
    <h5>Ingredients:</h5>
    <ul>
        ${ingredients.map(ingredient => `<li>${ingredient}</li>`).join('')}
    </ul>
    </div>
</div>`

const VideoHtml = `${meal.strYoutube ? `
<div class="row">
    <h5>Video Recipe</h5>
    <div class="videoWrapper">
        <iframe width="620" height="315"
        src="https://www.youtube.com/embed/${meal.strYoutube.slice(-11)}">
        </iframe>
    </div>
</div>` : ''}
`

    Selectors.header.innerHTML = HeaderMealSection
    Selectors.MealIngredient.innerHTML = MealIngredient
    Selectors.vidContainer.innerHTML = VideoHtml

}

async function getMealInfo() {
    try {
        const req = await fetch(RandomMealApi, {
            method: "GET"
        })
        const res = await req.json()
        console.log(res.meals[0])
        return res.meals[0]
    } catch (err) {
        throw new Error(err)
    }
}