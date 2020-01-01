using System.Collections;
using System.Collections.Generic;
using UnityEditor;
using UnityEngine;

// Custom Editor Script that will allow us to select either Level or World in the Editor


[CustomEditor(typeof(Level_Selection))]
public class Level_SelectionEditor : Editor
{
    string[] _choices = new[] { "Level", "World" };
    int _choiceIndex = 0;

    override public void OnInspectorGUI()
    {

        // Draw the default inspector
        var mc = target as Level_Selection;

        EditorGUILayout.PropertyField(serializedObject.FindProperty("ButtonUI"), true);
        EditorGUILayout.PropertyField(serializedObject.FindProperty("Content"), true);

        /*EditorGUILayout.PrefixLabel("Type");
        EditorGUI.indentLevel++;
        _choiceIndex = EditorGUILayout.Popup(_choiceIndex, _choices);*/




        EditorGUILayout.PropertyField(serializedObject.FindProperty("type"), true);
        _choiceIndex = serializedObject.FindProperty("type").enumValueIndex;
        EditorGUI.indentLevel++;


        //updated in code
        if (_choices[_choiceIndex] == "Level")
        {
            mc.type = Level_Selection.Type.Level;
            EditorGUILayout.PropertyField(serializedObject.FindProperty("LevelList"), true);
        }
        else
        {
            mc.type = Level_Selection.Type.World;
            EditorGUILayout.PropertyField(serializedObject.FindProperty("WorldList"), true);
        }

        serializedObject.ApplyModifiedProperties();


        // Save the changes back to the object
        EditorUtility.SetDirty(target);

    }


}


