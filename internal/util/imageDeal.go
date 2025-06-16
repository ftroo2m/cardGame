package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type ImageGenerationRequest struct {
	Model             string  `json:"model"`
	Prompt            string  `json:"prompt"`
	NegativePrompt    string  `json:"negative_prompt"`
	ImageSize         string  `json:"image_size"`
	BatchSize         int     `json:"batch_size"`
	Seed              int64   `json:"seed"`
	NumInferenceSteps int     `json:"num_inference_steps"`
	GuidanceScale     float64 `json:"guidance_scale"`
	PromptEnhancement bool    `json:"prompt_enhancement"`
}

type ImageGenerationResponse struct {
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
}

func ImageTobase64(name string) string {
	imagePath := "image/" + strings.Replace(name, "(S)", "", -1) + ".jpeg"
	file, err := os.Open(imagePath)
	if err != nil {
		log.Printf("Error: %v", err)
		return "error"
	}
	defer file.Close()
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error: %v", err)
		return "error"
	}
	base64Image := base64.StdEncoding.EncodeToString(imageData)
	return base64Image
}

func ImageToBase64Player(id int) string {
	imagePath := "player/" + strconv.Itoa(id) + ".jpeg"
	file, err := os.Open(imagePath)
	if err != nil {
		log.Printf("Error: %v", err)
		return "error"
	}
	defer file.Close()
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error: %v", err)
		return "error"
	}
	base64Image := base64.StdEncoding.EncodeToString(imageData)
	return base64Image
}

func NewPlayer(id int) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	description := []string{"Tall", "Lofty", "Slim", "Broad", "Round", "Square", "Flat", "Curved", "Steep", "Narrow", "Wide", "Long", "Short", "Tiny", "Massive", "Huge", "Petite", "Slender", "Robust", "Sturdy", "Angular", "Graceful", "Elegant", "Bulky", "Compact", "Lean", "Pointy", "Shapely", "Symmetrical", "Asymmetrical"}
	occupation := []string{"Knight", "Wizard", "Archer", "Warrior", "Thief", "Mage", "Cleric", "Bard", "Paladin", "Rogue", "Assassin", "Monk", "Druid", "Sorcerer", "Necromancer", "Hunter", "Alchemist", "Engineer", "Priest", "Shaman", "Blacksmith", "Merchant", "Scribe", "Carpenter", "Healer", "Sailor", "Fisherman", "Miner", "Farmer", "Cook"}

	combination := "Dungeons and Dragons" + description[rng.Intn(len(description))] + " " + description[rng.Intn(len(description))] + " " + description[rng.Intn(len(description))] + " " + occupation[rng.Intn(len(occupation))]
	seed := rng.Int31()
	url := "https://api.siliconflow.cn/v1/images/generations"

	// Create the request payload
	payload := ImageGenerationRequest{
		Model:             "Kwai-Kolors/Kolors",
		Prompt:            combination,
		NegativePrompt:    "",
		ImageSize:         "576x1024",
		BatchSize:         1,
		Seed:              int64(seed),
		NumInferenceSteps: 20,
		GuidanceScale:     7.5,
		PromptEnhancement: false,
	}

	// Marshal the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal payload: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Add headers
	req.Header.Add("Authorization", "your api token")
	req.Header.Add("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer res.Body.Close()

	// 解析响应数据
	var response ImageGenerationResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	if len(response.Images) > 0 {
		imageURL := response.Images[0].URL
		fmt.Printf("Downloading image: %s\n", imageURL)
		// 下载图片
		downloadImage(imageURL, "player/"+strconv.Itoa(id)+".jpeg")
	} else {
		log.Fatal("No image URL found in the response.")
	}
}

func downloadImage(url, filename string) {
	// 获取图片数据
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to download image: %v", err)
	}
	defer resp.Body.Close()

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// 保存图片到文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}

	fmt.Printf("Image saved as %s\n", filename)
}
