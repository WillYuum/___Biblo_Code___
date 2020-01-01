using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using UnityEngine.SceneManagement;
using UnityEditor;


public class Level_Selection : MonoBehaviour {

    //Choice of Type as either Level or World
    [HideInInspector]
    public enum Type {Level, World}
    [HideInInspector]
    public Type type;

    //Levels Class
    [System.Serializable]
    public class Levels
    {
        
        public string Level_Name;
        public int unlocked;
        public bool isInteractable;

    }

    //Worlds Class
    [System.Serializable]
    public class Worlds
    {
       
        public string World_name;
        public int starsCollected;
    }

    //List of Levels and Worlds
    public List<Levels> LevelList;
    public List<Worlds> WorldList;

    public GameObject ButtonUI;
    public Transform Content;

    void Awake()
    {
        FillContent();
    }

    [System.NonSerialized]
    public int LevelAmount;

    void FillContent()
    {
        if (type == Type.Level)
        {
            
            //----------for Levels-----------
            foreach (var Level in LevelList)
            {
               
                GameObject newButton = Instantiate(ButtonUI);
                LevelButton button = newButton.GetComponent<LevelButton>();
                button.LevelText.text = Level.Level_Name;

                if (PlayerPrefs.GetInt(Level.Level_Name) == 0)
                {
                    Level.unlocked = 0;
                    Level.isInteractable = false;
                }
                else
                {
                    Level.unlocked = 1;
                    Level.isInteractable = true;
                }


                button.unlocked = Level.unlocked;
                button.GetComponent<Button>().interactable = Level.isInteractable;
                button.GetComponent<Button>().onClick.AddListener(() => LoadScene(button.LevelText.text));
                button.transform.SetParent(Content,false);


                int currentScore = PlayerPrefs.GetInt(Level.Level_Name + "_score");

                if (currentScore > 0)
                {
                    button.star1.SetActive(true);
                }
                if (currentScore > 9)
                {
                    button.star2.SetActive(true);
                }
                if (currentScore > 12)
                {
                    button.star3.SetActive(true);
                }
            }


            

        }
        else
        {
            //-----------for World----------------
            if (type == Type.World)
            {
                foreach (var World in WorldList)
                {
                    GameObject newWorld = Instantiate(ButtonUI);
                    WorldUIButton ex_World = newWorld.GetComponent<WorldUIButton>();
                    ex_World.GetComponent<Button>().onClick.AddListener(() => LoadScene(ex_World.World_Name.text));
                    newWorld.transform.SetParent(Content, false);
                   
                }
            }
        }
        



    }
  

    void LoadScene(string LevelName)
    {
        SceneManager.LoadScene(LevelName);
    }


}




