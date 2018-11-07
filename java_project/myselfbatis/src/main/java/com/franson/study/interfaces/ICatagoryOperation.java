package com.franson.study.interfaces;


import com.franson.study.model.Catagory;

public interface ICatagoryOperation {
    void insertCatagory(Catagory catagory);
    Catagory getCatagorybyID(int id);
}

