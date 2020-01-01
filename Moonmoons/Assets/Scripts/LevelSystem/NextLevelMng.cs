using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class NextLevelMng : MonoBehaviour {

    int score = 1000;

    Level_Selection lvlMng;
    int NoOfLvl;
    private int currentLevel;

    


	// Use this for initialization
	void Start () {
        PlayerPrefs.SetInt("Level2", 1);
        PlayerPrefs.SetInt("Level1_score", score);

        StartCoroutine(Time());
         
     
	}

    

    // Update is called once per frame
    void Update()
    {
        
    }


    IEnumerator Time()
    {
        yield return new WaitForSeconds(2f);
        Application.LoadLevel(3);
    }

    void CheckCurrentLevel()
    {
        for(int i = 0; i <NoOfLvl; i++)
        {
            if (Application.loadedLevelName == "Level" + i)
            {
                currentLevel = i;
                SaveMyGame();
                
            }
        }
    }

    

    void SaveMyGame()
    {
        int NextLevel = currentLevel + 1;
        if (NextLevel < NoOfLvl)
        {
            PlayerPrefs.SetInt("Level" + NextLevel.ToString(), 1);
            PlayerPrefs.SetInt("Level" + currentLevel.ToString() + "_score", score);
        }
        else
        {
            PlayerPrefs.SetInt("Level" + currentLevel.ToString() + "_score", score);
        }
    }

}
